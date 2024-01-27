import dayjs from 'dayjs';
import LocalizedFormat from 'dayjs/plugin/localizedFormat';
import ja from 'dayjs/locale/ja';

dayjs.extend(LocalizedFormat)
dayjs.locale(ja);

export { dayjs };