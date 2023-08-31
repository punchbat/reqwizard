function getSecretkey() {
    const secretkeyMetaTag = document.querySelector('meta[name="secretkey"]');

    return secretkeyMetaTag ? secretkeyMetaTag.getAttribute("content") : null;
}

export { getSecretkey };
