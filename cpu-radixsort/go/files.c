

int *readfile(char *name)
{
    int ret, file_descriptor;
    if ((file_descriptor = open(name, O_RDONLY)) < 0)
        perror("open() error");
    else
    {
        ret = pread(file_descriptor, buf, ((sizeof(buf) - 1) - off), off);
        buf[ret] = 0x00;
        printf("block pread: \n<%s>\n", buf);
        if (close(file_descriptor) != 0)
            perror("close() error");
    }
    if (unlink("test.output") != 0)
        perror("unlink() error");
}