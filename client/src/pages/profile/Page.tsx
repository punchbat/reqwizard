/* eslint-disable jsx-a11y/anchor-is-valid */
import { useState, useCallback, FC, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Typography, Spin, Button, Input, Form, Upload, Select, DatePicker } from "antd";
import { SaveOutlined, EditOutlined, UploadOutlined, DeleteOutlined } from "@ant-design/icons";
import { RcFile, UploadChangeParam, UploadFile } from "antd/es/upload";

import { cn, getBase64 } from "@utils";
import {
    useGetMyProfileQuery,
    useGetProfileQuery,
    useGetRolesQuery,
    useLogoutMutation,
    useUpdateProfileMutation,
} from "@app/services/auth";
import { useAppDispatch } from "@hooks/store-hooks";
import { ToastStore } from "@widgets/index";
import { EUserGender, FailResponse, IUser } from "@localtypes";

import "./Page.scss";
import dayjs, { Dayjs } from "dayjs";
import "dayjs/locale/ru"; // Импорт нужной локали

dayjs.locale("ru");

const { Text } = Typography;
const { Option } = Select;

type FormValues = Omit<IUser, "_id" | "email" | "verified" | "applicationCreatedAt" | "createdAt" | "updatedAt"> & {
    avatar: UploadChangeParam<UploadFile>;
    userRoles: Array<string>;
    birthday: string & Dayjs;
};

const b = cn("profile");

const Profile: FC = function () {
    const { id } = useParams();
    const navigate = useNavigate();

    const dispatch = useAppDispatch();
    const { data: dataRoles, isLoading: isLoadingRoles } = useGetRolesQuery();

    const [updateProfile, { isLoading: isUpdateProfileLoading }] = useUpdateProfileMutation();
    const { data: myProfileData } = useGetMyProfileQuery();
    const { data, isLoading } = useGetProfileQuery(id || "");
    const [logout] = useLogoutMutation();

    const [isEditMode, setIsEditMode] = useState<boolean>(false);
    const [avatarContent, setAvatarContent] = useState<string>();

    const [form] = Form.useForm<FormValues>();
    const setFormDefaultValues = useCallback(() => {
        form.setFieldsValue({
            avatar: data?.payload?.avatar || "",
            name: data?.payload?.name || "",
            surname: data?.payload?.surname || "",
            gender: data?.payload?.gender || EUserGender.MALE,
            userRoles: data?.payload?.userRoles.map(i => i.name),
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore
            birthday: data?.payload?.birthday ? dayjs(data?.payload?.birthday) : undefined,
        });

        setAvatarContent(
            data?.payload.avatar && data?.payload.avatar.length
                ? `${import.meta.env.VITE_API_URL}/uploads/avatars/${data?.payload.avatar}`
                : `/assets/user/${data?.payload.gender}.png`,
        );
    }, [data?.payload, form]);

    useEffect(() => {
        setFormDefaultValues();
    }, [data?.payload, form, setFormDefaultValues]);

    const isMyProfile = () => {
        // eslint-disable-next-line no-underscore-dangle
        return id === myProfileData?.payload._id;
    };

    const onFinish = async (values: FormValues) => {
        try {
            const formData = new FormData();
            values.userRoles?.forEach((name: string) => {
                formData.append("userRoles", name);
            });
            formData.append("name", values.name);
            formData.append("surname", values.surname);
            formData.append("gender", values.gender);
            formData.append("birthday", dayjs(values.birthday).toISOString());

            if (values.avatar && values.avatar.file) {
                const content = await getBase64(values.avatar.file as RcFile);
                if (content) {
                    const blob = await fetch(content).then(res => res.blob());
                    formData.append("avatar", blob, values.avatar.file.name);
                }
            }

            await updateProfile(formData).unwrap();
        } catch (err) {
            const { data } = err as { data: FailResponse };
            dispatch(
                ToastStore.notify({
                    // eslint-disable-next-line @typescript-eslint/no-unsafe-member-access, @typescript-eslint/no-unsafe-assignment
                    message: data?.message,
                    options: {
                        type: ToastStore.MessageType.ERROR,
                        duration: 3000,
                        position: ToastStore.MessagePositions["RIGHT-TOP"],
                    },
                }),
            );
        }
    };

    const handleExit = async () => {
        try {
            await logout().unwrap();

            navigate("/sign-in");
        } catch (err) {
            dispatch(
                ToastStore.notify({
                    message: (err as Object).toString(),
                    options: {
                        type: ToastStore.MessageType.ERROR,
                        duration: 3000,
                        position: ToastStore.MessagePositions["RIGHT-TOP"],
                    },
                }),
            );
        }
    };

    if (isLoading || isUpdateProfileLoading || isLoadingRoles) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <Form onFinish={onFinish} form={form}>
                    <div className={b("top")}>
                        <div className={b("picture")}>
                            {isEditMode ? (
                                <Form.Item
                                    name="avatar"
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
                                        className={b("upload")}
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
                                        <div className={b("upload_wrapper")}>
                                            {avatarContent ? (
                                                <>
                                                    <img
                                                        className={b("upload_img")}
                                                        src={avatarContent}
                                                        alt="Avatar"
                                                        style={{ width: "100%" }}
                                                    />
                                                    <div className={b("upload_desc")}>
                                                        <UploadOutlined />
                                                        <Text>New Avatar</Text>
                                                    </div>
                                                </>
                                            ) : (
                                                <>
                                                    <UploadOutlined />
                                                    <Text>New Avatar</Text>
                                                </>
                                            )}
                                        </div>
                                    </Upload>
                                </Form.Item>
                            ) : (
                                <img src={avatarContent} alt="" />
                            )}
                        </div>

                        <div className={b("info")}>
                            <div className={b("email")}>
                                <Text>{data?.payload?.email}</Text>
                            </div>

                            <div className={b("fullname")}>
                                <div className={b("surname")}>
                                    {isEditMode ? (
                                        <Form.Item
                                            name="surname"
                                            rules={[{ required: true, message: "Please enter your surname!" }]}
                                        >
                                            <Input placeholder="Surname" />
                                        </Form.Item>
                                    ) : (
                                        <Text>{form.getFieldValue("surname")}</Text>
                                    )}
                                </div>
                                <div className={b("name")}>
                                    {isEditMode ? (
                                        <Form.Item
                                            name="name"
                                            rules={[{ required: true, message: "Please enter your name!" }]}
                                        >
                                            <Input placeholder="Name" />
                                        </Form.Item>
                                    ) : (
                                        <Text>{form.getFieldValue("name")}</Text>
                                    )}
                                </div>
                            </div>

                            <div className={b("gender")}>
                                {isEditMode ? (
                                    <Form.Item
                                        name="gender"
                                        rules={[{ required: true, message: "Please select your gender!" }]}
                                    >
                                        <Select placeholder="Select gender">
                                            {Object.entries(EUserGender).map(([key, value]) => (
                                                <Option key={key} value={value}>
                                                    {value}
                                                </Option>
                                            ))}
                                        </Select>
                                    </Form.Item>
                                ) : (
                                    <Text>{form.getFieldValue("gender")}</Text>
                                )}
                            </div>
                            <div className={b("birthday")}>
                                {isEditMode ? (
                                    <Form.Item
                                        name="birthday"
                                        rules={[{ required: true, message: "Please select your birthday!" }]}
                                        wrapperCol={{ span: 24 }}
                                    >
                                        <DatePicker
                                            placeholder="Select birthday"
                                            style={{ width: "100%" }}
                                            showTime={false}
                                        />
                                    </Form.Item>
                                ) : (
                                    <Text>
                                        {form.getFieldValue("birthday")
                                            ? new Date(form.getFieldValue("birthday") as string).toLocaleDateString(
                                                  "en-US",
                                                  {
                                                      year: "numeric",
                                                      month: "long",
                                                      day: "numeric",
                                                  },
                                              )
                                            : "Not found birthday"}
                                    </Text>
                                )}
                            </div>
                        </div>
                    </div>

                    <div className={b("roles")}>
                        {isEditMode ? (
                            <Form.Item
                                name="userRoles"
                                rules={[{ required: true, message: "Please select your gender!" }]}
                            >
                                <Select
                                    mode="multiple"
                                    style={{ width: "100%" }}
                                    placeholder="Select role"
                                    optionLabelProp="label"
                                >
                                    {dataRoles?.payload.map(i => (
                                        <Option key={i.name} value={i.name} disabled={i.name === "user"}>
                                            {i.name}
                                        </Option>
                                    ))}
                                </Select>
                            </Form.Item>
                        ) : (
                            form.getFieldValue("userRoles")?.map((name: string) => (
                                // eslint-disable-next-line react/no-array-index-key
                                <div className={b("role")} key={name}>
                                    <Text>{name}</Text>
                                </div>
                            ))
                        )}
                    </div>
                    <div className={b("additional_info")}>
                        <div className={b("last_application_time")}>
                            <Text style={{ fontSize: "10px" }}>Last Application is created at: </Text>
                            <Text style={{ fontSize: "10px" }} strong>
                                {data?.payload?.applicationCreatedAt
                                    ? new Date(data?.payload?.applicationCreatedAt).toLocaleDateString("en-US", {
                                          year: "numeric",
                                          month: "long",
                                          day: "numeric",
                                          hour: "2-digit",
                                          minute: "2-digit",
                                          second: "2-digit",
                                          timeZoneName: "short",
                                      })
                                    : "Not found"}
                            </Text>
                        </div>
                        <div className={b("created_at")}>
                            <Text style={{ fontSize: "10px" }}>Account is created at: </Text>
                            <Text style={{ fontSize: "10px" }} strong>
                                {data?.payload?.createdAt
                                    ? new Date(data?.payload?.createdAt).toLocaleDateString("en-US", {
                                          year: "numeric",
                                          month: "long",
                                          day: "numeric",
                                          hour: "2-digit",
                                          minute: "2-digit",
                                          second: "2-digit",
                                          timeZoneName: "short",
                                      })
                                    : "Not found"}
                            </Text>
                        </div>
                    </div>
                    <div className={b("actions")}>
                        <Button type="link" htmlType="button" onClick={handleExit}>
                            Exit
                        </Button>
                    </div>
                </Form>
            </div>
            {isMyProfile() && (
                <div className={b("edit_action")}>
                    <Button
                        type={isEditMode ? "primary" : "default"}
                        htmlType="button"
                        onClick={() => {
                            setIsEditMode(!isEditMode);

                            if (isEditMode) {
                                form.submit();
                            }
                        }}
                        icon={isEditMode ? <SaveOutlined /> : <EditOutlined />}
                    >
                        {isEditMode ? "Save" : "Edit"}
                    </Button>
                    {isEditMode && (
                        <Button
                            htmlType="button"
                            onClick={() => {
                                setFormDefaultValues();

                                setIsEditMode(!isEditMode);
                            }}
                            icon={<DeleteOutlined />}
                        >
                            Cancel
                        </Button>
                    )}
                </div>
            )}
        </div>
    );
};

export { Profile };
