import axios from 'axios';

const RESERVATION_API_BASE_URL = "/Reservation";

class ReservationService {

    getReservations(){
        return axios.get(RESERVATION_API_BASE_URL + '/' );
    }

    createReservation(reservation){
        return axios.post(RESERVATION_API_BASE_URL  + '/create', reservation);
    }

    getReservationById(reservationId){
        return axios.get(RESERVATION_API_BASE_URL + '/load?reservationId=' + reservationId);
    }

    updateReservation(reservation){
        return axios.put(RESERVATION_API_BASE_URL + '/update', reservation);
    }

    deleteReservation(reservationId){
        return axios.delete(RESERVATION_API_BASE_URL + '/delete?reservationId=' + reservationId);
    }
}

export default new ReservationService()