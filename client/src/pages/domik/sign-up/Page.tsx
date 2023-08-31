import { FC, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Form, Input, Typography, Select, Spin } from "antd";
import { LockOutlined, UserOutlined } from "@ant-design/icons";
import { cn } from "@utils";

import { FailResponse, ResponseInterface, isSuccessResponse } from "@localtypes";
import {
    SignUpInput,
    CheckVerifyCodeInput,
    useSignUpMutation,
    useSendVerifyCodeMutation,
    useCheckVerifyCodeMutation,
    useGetRolesQuery,
} from "@app/services/auth";
import { useAppDispatch } from "@hooks/index";
import { ToastStore } from "@widgets/index";

import "./Page.scss";

const { Title } = Typography;
const { Option } = Select;

const b = cn("signin");

const SignUp: FC = function () {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const { data: dataRoles, isLoading: isLoadingRoles } = useGetRolesQuery();
    const [signUp, { isLoading: isSignUpLoading }] = useSignUpMutation();
    const [sendVerifyCode, { isLoading: isSendVerifyCodeLoading }] = useSendVerifyCodeMutation();
    const [checkVerifyCode, { isLoading: isCheckVerifyCodeLoading }] = useCheckVerifyCodeMutation();

    const [isCheckMode, setIsCheckMode] = useState(false);
    const [userInput, setUserInput] = useState<SignUpInput>();

    const onFinish = async (values: SignUpInput | CheckVerifyCodeInput) => {
        if (!isCheckMode) {
            setUserInput(values as SignUpInput);

            try {
                await signUp(values as SignUpInput).unwrap();

                await sendVerifyCode(values).unwrap();

                setIsCheckMode(true);
            } catch (err) {
                const { data } = err as { data: FailResponse };
                dispatch(
                    ToastStore.notify({
                        message: data?.message,
                        options: {
                            type: ToastStore.MessageType.ERROR,
                            duration: 3000,
                            position: ToastStore.MessagePositions["RIGHT-TOP"],
                        },
                    }),
                );
            }
        } else {
            try {
                const response: ResponseInterface<string> = await checkVerifyCode({
                    email: userInput?.email || "",
                    password: userInput?.password || "",
                    verifyCode: (values as CheckVerifyCodeInput).verifyCode,
                }).unwrap();

                if (!response || isSuccessResponse<string>(response)) {
                    navigate("/");
                }
            } catch (err) {
                const { data } = err as { data: FailResponse };
                dispatch(
                    ToastStore.notify({
                        message: data?.message,
                        options: {
                            type: ToastStore.MessageType.ERROR,
                            duration: 3000,
                            position: ToastStore.MessagePositions["RIGHT-TOP"],
                        },
                    }),
                );
            }
        }
    };

    const handleSendVerifyCode = async () => {
        await sendVerifyCode(userInput as SignUpInput).unwrap();
    };

    const handleBack = () => {
        navigate("/sign-in");
    };

    if (isLoadingRoles || isSignUpLoading || isSendVerifyCodeLoading || isCheckVerifyCodeLoading) {
        return (
            <div className={b("spin")}>
                <Spin />
            </div>
        );
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <Title level={1}>{!isCheckMode ? "Sign-up" : "Verify code"}</Title>
                </div>
                <div className={b("content")}>
                    <Form name="basic" initialValues={{ remember: true }} onFinish={onFinish} layout="vertical">
                        {!isCheckMode ? (
                            <>
                                <Form.Item name="role" rules={[{ required: true, message: "Please select role!" }]}>
                                    <Select placeholder="select your role">
                                        {dataRoles?.payload?.map(i => (
                                            <Option value={i.name} key={i.name}>
                                                {i.name}
                                            </Option>
                                        ))}
                                    </Select>
                                </Form.Item>

                                <Form.Item
                                    name="email"
                                    rules={[
                                        {
                                            required: true,
                                            type: "email",
                                            message: "The input is not valid Email!",
                                        },
                                    ]}
                                >
                                    <Input
                                        prefix={<UserOutlined className="site-form-item-icon" />}
                                        placeholder="Email"
                                    />
                                </Form.Item>

                                <Form.Item
                                    name="password"
                                    rules={[
                                        { required: true, message: "Please input your password!" },
                                        () => ({
                                            validator(_, value) {
                                                const password = value as string;

                                                const hasUpperCase = /[A-Z]/.test(password);
                                                const hasLowerCase = /[a-z]/.test(password);
                                                const hasNumber = /\d/.test(password);
                                                const hasSymbol = /[!@#$%^&*()_+\-=\\[\]{};':"\\|,.<>\\/?]/.test(
                                                    password,
                                                );
                                                const isLengthValid = password.length >= 8;

                                                if (
                                                    hasUpperCase &&
                                                    hasLowerCase &&
                                                    hasNumber &&
                                                    hasSymbol &&
                                                    isLengthValid
                                                ) {
                                                    return Promise.resolve();
                                                }

                                                return Promise.reject(
                                                    new Error(
                                                        "The password must contain at least one uppercase letter, one lowercase letter, one digit, one symbol, and be at least 8 characters long!",
                                                    ),
                                                );
                                            },
                                        }),
                                    ]}
                                >
                                    <Input
                                        prefix={<LockOutlined className="site-form-item-icon" />}
                                        type="password"
                                        placeholder="Password"
                                    />
                                </Form.Item>

                                <Form.Item
                                    name="passwordConfirm"
                                    dependencies={["password"]}
                                    hasFeedback
                                    rules={[
                                        {
                                            required: true,
                                            message: "Please confirm your password!",
                                        },
                                        ({ getFieldValue }) => ({
                                            validator(_, value) {
                                                if (!value || getFieldValue("password") === value) {
                                                    return Promise.resolve();
                                                }
                                                return Promise.reject(
                                                    new Error("The new password that you entered do not match!"),
                                                );
                                            },
                                        }),
                                    ]}
                                >
                                    <Input
                                        prefix={<LockOutlined className="site-form-item-icon" />}
                                        type="password"
                                        placeholder="Confirm Password"
                                    />
                                </Form.Item>
                            </>
                        ) : (
                            <Form.Item
                                name="verifyCode"
                                hasFeedback
                                rules={[
                                    {
                                        required: true,
                                        message: "Please enter the verification code!",
                                    },
                                    {
                                        pattern: /^\d{6}$/,
                                        message: "The verification code must consist of 6 digits!",
                                    },
                                ]}
                            >
                                <Input
                                    prefix={<LockOutlined className="site-form-item-icon" />}
                                    placeholder="Verification Code"
                                    min={0}
                                    max={999999}
                                    onInput={e => {
                                        const inputValue = e.currentTarget.value;
                                        if (inputValue.length > 6) {
                                            e.currentTarget.value = inputValue.slice(0, 6);
                                        }
                                    }}
                                />
                            </Form.Item>
                        )}

                        <Form.Item>
                            <Button
                                type="primary"
                                htmlType="submit"
                                loading={isSignUpLoading || isSendVerifyCodeLoading || isCheckVerifyCodeLoading}
                            >
                                {!isCheckMode ? "Sign-up" : "Verify code"}
                            </Button>
                            {isCheckMode && (
                                <Button
                                    type="text"
                                    htmlType="button"
                                    onClick={handleSendVerifyCode}
                                    style={{ marginLeft: "4px" }}
                                >
                                    Send verify code again
                                </Button>
                            )}
                        </Form.Item>
                        <Form.Item>
                            <Button type="link" onClick={handleBack}>
                                Back to sign-in
                            </Button>
                        </Form.Item>
                    </Form>
                </div>
            </div>
        </div>
    );
};

export { SignUp };
