import { useState, FC } from "react";
import { useNavigate } from "react-router-dom";
import { Button, Form, Input, Typography, Select, Spin, Upload, DatePicker } from "antd";
import { LockOutlined, UserOutlined, UploadOutlined } from "@ant-design/icons";
import { cn, getBase64 } from "@utils";

import { FailResponse, ResponseInterface, isSuccessResponse } from "@localtypes";
import {
    useSignUpMutation,
    useSendVerifyCodeMutation,
    useCheckVerifyCodeMutation,
    useGetRolesQuery,
} from "@app/services/auth";
import { useAppDispatch } from "@hooks/index";
import { ToastStore } from "@widgets/index";
import { RcFile, UploadChangeParam, UploadFile } from "antd/es/upload";
import dayjs from "dayjs";

import "./Page.scss";

const { Title, Text } = Typography;
const { Option } = Select;

const b = cn("signin");

export interface FormValues {
    avatar: UploadChangeParam<UploadFile>;
    name: string;
    surname: string;
    gender: string;
    birthday: string;
    email: string;
    password: string;
    passwordConfirm: string;
    role: string;

    // check verify code
    verifyCode: string;
}

const SignUp: FC = function () {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    const { data: dataRoles, isLoading: isLoadingRoles } = useGetRolesQuery();
    const [signUp, { isLoading: isSignUpLoading }] = useSignUpMutation();
    const [sendVerifyCode, { isLoading: isSendVerifyCodeLoading }] = useSendVerifyCodeMutation();
    const [checkVerifyCode, { isLoading: isCheckVerifyCodeLoading }] = useCheckVerifyCodeMutation();

    const [isCheckMode, setIsCheckMode] = useState(false);
    const [userInput, setUserInput] = useState<FormValues>();
    const [avatarContent, setAvatarContent] = useState<string>();

    const onFinish = async (values: FormValues) => {
        if (!isCheckMode) {
            setUserInput(values);

            try {
                const localValues: FormValues = values;

                const formData = new FormData();
                formData.append("role", localValues.role);
                formData.append("email", localValues.email);
                formData.append("password", localValues.password);
                formData.append("passwordConfirm", localValues.passwordConfirm);
                formData.append("name", localValues.name);
                formData.append("surname", localValues.surname);
                formData.append("gender", localValues.gender);
                formData.append("birthday", dayjs(localValues.birthday).toISOString());

                if (localValues?.avatar) {
                    const content = await getBase64(localValues?.avatar?.file as RcFile);

                    if (content) {
                        const blob = await fetch(content).then(res => res.blob());
                        formData.append("avatar", blob, localValues.avatar.file.name);
                    }
                }

                await signUp(formData).unwrap();

                await sendVerifyCode(localValues).unwrap();

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
                    verifyCode: values.verifyCode,
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
        await sendVerifyCode(userInput as FormValues).unwrap();
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
            <div className={b("wrapper")}>
                <div className={b("inner")}>
                    <div className={b("title")}>
                        <Title level={1}>{!isCheckMode ? "Sign-up" : "Verify code"}</Title>
                    </div>
                    <div className={b("content")}>
                        <Form
                            name="basic"
                            initialValues={{ remember: true }}
                            onFinish={onFinish}
                            layout="vertical"
                            size="large"
                        >
                            {!isCheckMode ? (
                                <>
                                    <Form.Item
                                        name="avatar"
                                        wrapperCol={{
                                            span: 6,
                                            offset: 9,
                                        }}
                                        rules={[
                                            {
                                                validator: async (_, value) => {
                                                    if (!value) {
                                                        return Promise.resolve();
                                                    }

                                                    const isImage = value.file?.type?.startsWith("image/");
                                                    const isSizeValid =
                                                        value.file?.size && value.file?.size <= 2 * 1024 * 1024; // 2 MB in bytes

                                                    if (!isImage) {
                                                        return Promise.reject(
                                                            new Error("You can only upload image files!"),
                                                        );
                                                    }

                                                    if (!isSizeValid) {
                                                        return Promise.reject(
                                                            new Error("Image size should not exceed 2MB!"),
                                                        );
                                                    }

                                                    return Promise.resolve();
                                                },
                                            },
                                        ]}
                                    >
                                        <Upload
                                            name="avatar"
                                            listType="picture-card"
                                            maxCount={1}
                                            beforeUpload={() => {
                                                return false;
                                            }}
                                            onChange={async (info: UploadChangeParam<UploadFile<any>>) => {
                                                const content = await getBase64(info.file as RcFile);
                                                setAvatarContent(content);
                                            }}
                                            showUploadList={false}
                                        >
                                            <div className={b("upload")}>
                                                {avatarContent ? (
                                                    <img src={avatarContent} alt="Avatar" style={{ width: "100%" }} />
                                                ) : (
                                                    <>
                                                        <UploadOutlined />
                                                        <Text>Your Avatar</Text>
                                                    </>
                                                )}
                                            </div>
                                        </Upload>
                                    </Form.Item>

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

                                    <Form.Item
                                        name="birthday"
                                        rules={[{ required: true, message: "Please select your birthday!" }]}
                                        wrapperCol={{ span: 24 }}
                                    >
                                        <DatePicker placeholder="Select birthday" style={{ width: "100%" }} />
                                    </Form.Item>
                                    <Form.Item
                                        name="name"
                                        rules={[{ required: true, message: "Please enter your name!" }]}
                                    >
                                        <Input placeholder="Name" />
                                    </Form.Item>
                                    <Form.Item
                                        name="surname"
                                        rules={[{ required: true, message: "Please enter your surname!" }]}
                                    >
                                        <Input placeholder="Surname" />
                                    </Form.Item>
                                    <Form.Item
                                        name="gender"
                                        rules={[{ required: true, message: "Please select your gender!" }]}
                                    >
                                        <Select placeholder="Select gender">
                                            <Option value="male">Male</Option>
                                            <Option value="female">Female</Option>
                                            <Option value="other">Other</Option>
                                        </Select>
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
                                    <Button type="link" onClick={handleSendVerifyCode}>
                                        Send verify code again
                                    </Button>
                                )}
                                {!isCheckMode && (
                                    <Button type="link" onClick={handleBack}>
                                        Back to sign-in
                                    </Button>
                                )}
                            </Form.Item>
                        </Form>
                    </div>
                </div>
            </div>
        </div>
    );
};

export { SignUp };
