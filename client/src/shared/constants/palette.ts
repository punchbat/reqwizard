export interface IPalette {
    /**
     * @name BorderRadius
     * @default 8
     */
    borderRadius: number;
    /**
     * @name Brand Color
     * @desc Brand color is one of the most direct visual elements to reflect the characteristics and communication of the product. After you have selected the brand color, we will automatically generate a complete color palette and assign it effective design semantics.
     */
    colorPrimary: string;
}

export const PALETTE: IPalette = {
    borderRadius: 8,
    colorPrimary: "#3D5AFE",
};

export const applicationStatusColors = {
    canceled: "red",
    waiting: "orange",
    working: "blue",
    done: "green",
};
