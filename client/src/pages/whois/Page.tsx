import { FC, useState, useEffect } from "react";
import { Typography } from "antd";
import FingerprintJS, { GetResult } from "@fingerprintjs/fingerprintjs";

import { cn } from "@utils";

import "./Page.scss";

const { Title, Text } = Typography;

const b = cn("whois");

const Whois: FC = function () {
    const fingerprint = FingerprintJS.load();

    const [data, setData] = useState<GetResult>();

    useEffect(() => {
        fingerprint
            .then(fp => fp.get())
            .then(result => {
                setData(result);
            });
    });

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("title")}>
                    <Title level={1}>Who are you?</Title>
                </div>
                <div className={b("content")}>
                    {Array.isArray(data?.components.languages) ? (
                        <div className={b("langs")}>
                            <Text>Do you speak these languages?</Text>
                            {data?.components?.languages?.map((lang: Array<string>) => (
                                <div className={b("lang")} key={lang[0]}>
                                    <Text>{lang[0]}</Text>
                                </div>
                            ))}
                        </div>
                    ) : null}
                    {data?.components.platform.value ? (
                        <div className={b("platform")}>
                            <Text>
                                Are you using a <strong>{data?.components.platform.value}</strong> laptop?
                            </Text>
                        </div>
                    ) : null}
                    {data?.components.timezone.value ? (
                        <div className={b("timezone")}>
                            <Text>
                                Are you from <strong>{data?.components.timezone.value}</strong>?
                            </Text>
                        </div>
                    ) : null}
                    {data?.components.vendorFlavors.value ? (
                        <div className={b("vendor")}>
                            <Text>
                                Are you using <strong>{data?.components.vendorFlavors.value}</strong> brower?
                            </Text>
                        </div>
                    ) : null}
                </div>
            </div>
        </div>
    );
};

export { Whois };
