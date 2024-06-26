export const getImgPath = (val1, val2) => {

    if (!val1 || !val2) {
        return '';
    }

    let name = val1;
    let ext = val2;

    if (ext[0] !== '.') {
        ext = `.${ext}`
    }

    const path = `/resources/${ext}/${name}${ext}`
    return path;
} 