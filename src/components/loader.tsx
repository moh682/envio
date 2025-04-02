type Props = {
  text?: string;
};

export const Loader = ({ text }: Props) => {
  return (
    <div className="flex items-center gap-2 ">
      <div
        className="animate-spin inline-block size-4 border-2 border-current border-t-transparent text-white rounded-full dark:text-blue-500"
        role="status"
        aria-label="loading"
      >
        <span className="sr-only">Loading...</span>
      </div>
      {text && <span>{text}</span>}
    </div>
  );
};
