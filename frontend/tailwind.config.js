/** @type {import('tailwindcss').Config} */
module.exports = {
    corePlugins: {
        preflight: false
    },
    content: ["./index.html", "./src/**/*.{vue,js}"],
    theme: {
        extend: {
            colors: {
                bg_color: "var(--el-bg-color)",
                primary: "var(--el-color-primary)",
                text_color_primary: "var(--el-text-color-primary)",
                text_color_regular: "var(--el-text-color-regular)"
            }
        },
    },
    plugins: [],
}