import axios from 'axios';

const CHANNEL_API_BASE_URL = "/Channel";

class ChannelService {

    getChannels(){
        return axios.get(CHANNEL_API_BASE_URL + '/' );
    }

    createChannel(channel){
        return axios.post(CHANNEL_API_BASE_URL  + '/create', channel);
    }

    getChannelById(channelId){
        return axios.get(CHANNEL_API_BASE_URL + '/load?channelId=' + channelId);
    }

    updateChannel(channel){
        return axios.put(CHANNEL_API_BASE_URL + '/update', channel);
    }

    deleteChannel(channelId){
        return axios.delete(CHANNEL_API_BASE_URL + '/delete?channelId=' + channelId);
    }
}

export default new ChannelService()