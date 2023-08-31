import { useState, useCallback, FC } from "react";
import { useLocation, useNavigate } from "react-router-dom";

import { SearchOutlined, UpOutlined, DownOutlined } from "@ant-design/icons";
import { Button, Input, Spin, Typography, DatePicker, TimeRangePickerProps } from "antd";
import { cn } from "@utils";
import { FilterInput, useGetTicketResponsesByManagerIDQuery } from "@app/services/ticket_response";
import { TicketResponseItem } from "@features/index";
import dayjs from "dayjs";
import qs from "qs";

import "./index.scss";

const { Text } = Typography;
const { RangePicker } = DatePicker;

const b = cn("ticket_responses");

interface Filter {
    search?: string;
    createdAtRange?: [string | undefined, string | undefined];
    updatedAtRange?: [string | undefined, string | undefined];
}

const rangePresets: TimeRangePickerProps["presets"] = [
    { label: "Now", value: [dayjs(), dayjs().endOf("day")] },
    { label: "Last 7 Days", value: [dayjs().add(-7, "d"), dayjs()] },
    { label: "Last 14 Days", value: [dayjs().add(-14, "d"), dayjs()] },
    { label: "Last 30 Days", value: [dayjs().add(-30, "d"), dayjs()] },
    { label: "Last 90 Days", value: [dayjs().add(-90, "d"), dayjs()] },
];

const TicketResponses: FC = function () {
    const navigate = useNavigate();
    const location = useLocation();

    const queryObj = qs.parse(location.search.substring(1));
    const { tab } = queryObj as {
        tab: string;
    };
    const { search, createdAtFrom, createdAtTo, updatedAtFrom, updatedAtTo } = queryObj as FilterInput;

    const { data, isLoading, refetch } = useGetTicketResponsesByManagerIDQuery({
        search,
        createdAtFrom: createdAtFrom && dayjs(createdAtFrom).toISOString(),
        createdAtTo: createdAtTo && dayjs(createdAtTo).toISOString(),
        updatedAtFrom: updatedAtFrom && dayjs(updatedAtFrom).toISOString(),
        updatedAtTo: updatedAtTo && dayjs(updatedAtTo).toISOString(),
    });

    const [filter, setFilter] = useState<Filter>({
        search,
        createdAtRange: [createdAtFrom, createdAtTo],
        updatedAtRange: [updatedAtFrom, updatedAtTo],
    });

    const [isVisibleSubFilter, setIsVisibleSubFilter] = useState<boolean>(
        !![createdAtFrom, createdAtTo, updatedAtFrom, updatedAtTo].filter(Boolean).length,
    );

    const [typingTimeout, setTypingTimeout] = useState<NodeJS.Timeout | null>(null);

    const handleSearchChange = (value: Filter["search"]) => {
        if (value && value.length) {
            setFilter(prevFilter => ({
                ...prevFilter,
                search: value,
            }));
        } else {
            setFilter(prevFilter => {
                delete prevFilter.search;

                return prevFilter;
            });
        }

        if (typingTimeout) {
            clearTimeout(typingTimeout);
        }

        setTypingTimeout(
            setTimeout(() => {
                // eslint-disable-next-line @typescript-eslint/no-use-before-define
                handleApplyFilterAndRefetch(value && value.length ? { search: value } : {});
            }, 500),
        );
    };

    const handleFiltersItemChange = <K extends keyof Filter>(key: K, value: Filter[K]) => {
        if (
            (key === "createdAtRange" || key === "updatedAtRange") &&
            !(Array.isArray(value) && value.filter(Boolean).length)
        ) {
            setFilter(prevFilter => {
                delete prevFilter[key];

                return prevFilter;
            });
            return;
        }

        setFilter(prevFilter => ({
            ...prevFilter,
            [key]: value,
        }));
    };

    const handleApplyFilterAndRefetch = useCallback(
        (concatFilter?: FilterInput) => {
            const localFilter = { ...filter, ...concatFilter };

            let createdAtFrom;
            let createdAtTo;
            if (localFilter.createdAtRange) {
                [createdAtFrom, createdAtTo] = localFilter.createdAtRange;
                delete localFilter.createdAtRange;
            }

            let updatedAtFrom;
            let updatedAtTo;
            if (localFilter.updatedAtRange) {
                [updatedAtFrom, updatedAtTo] = localFilter.updatedAtRange;
                delete localFilter.updatedAtRange;
            }

            const query: FilterInput = { ...localFilter, createdAtFrom, createdAtTo, updatedAtFrom, updatedAtTo };

            navigate({
                pathname: location.pathname,
                search: qs.stringify({ ...query, tab }),
            });

            refetch();
        },
        [filter, location.pathname, navigate, refetch, tab],
    );

    if (isLoading) {
        return <Spin />;
    }

    return (
        <div className={b()}>
            <div className={b("inner")}>
                <div className={b("filters")}>
                    <div className={b("main")}>
                        <Input
                            placeholder="Search"
                            value={filter.search || undefined}
                            onChange={e => handleSearchChange(e.target.value)}
                        />
                    </div>
                    <div className={b("sub", isVisibleSubFilter ? ["active"] : [])}>
                        <div className={b("sub_content")}>
                            <RangePicker
                                presets={rangePresets}
                                showTime
                                format="YYYY-MM-DD HH:mm:ss"
                                value={
                                    filter.createdAtRange
                                        ? [dayjs(filter.createdAtRange[0]), dayjs(filter.createdAtRange[1])]
                                        : undefined
                                }
                                onChange={(_, dateStrings) => handleFiltersItemChange("createdAtRange", dateStrings)}
                                placeholder={["Created From", "Created To"]}
                            />
                            <RangePicker
                                presets={rangePresets}
                                showTime
                                format="YYYY-MM-DD HH:mm:ss"
                                value={
                                    filter.updatedAtRange
                                        ? [dayjs(filter.updatedAtRange[0]), dayjs(filter.updatedAtRange[1])]
                                        : undefined
                                }
                                onChange={(_, dateStrings) => handleFiltersItemChange("updatedAtRange", dateStrings)}
                                placeholder={["Updated From", "Updated To"]}
                            />
                        </div>
                        <div className={b("sub_submit")}>
                            <Button icon={<SearchOutlined />} onClick={() => handleApplyFilterAndRefetch()}>
                                Search
                            </Button>
                        </div>
                    </div>
                    <div className={b("sub_show")}>
                        <Button
                            type="text"
                            onClick={() => setIsVisibleSubFilter(!isVisibleSubFilter)}
                            icon={isVisibleSubFilter ? <UpOutlined /> : <DownOutlined />}
                        >
                            Add more filters
                        </Button>
                    </div>
                </div>
                <div className={b("content")}>
                    {data?.payload?.length ? (
                        <div className={b("items")}>
                            {data?.payload?.map((i, index) => (
                                // eslint-disable-next-line react/no-array-index-key, react/jsx-props-no-spreading
                                <TicketResponseItem key={index} {...i} />
                            ))}
                        </div>
                    ) : (
                        <div>
                            <Text>No active applications</Text>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export { TicketResponses };
