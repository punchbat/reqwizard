import { RowModeType, RowType } from '@cshared/types/tables';

export const getCalcRowType = (rowMode: RowModeType, rowType: RowType) => {
    return rowMode === RowModeType.UPDATE && rowType === RowType.OLD ? RowType.UPDATED : rowType;
};
