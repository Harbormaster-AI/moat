import React, { Component } from 'react'
import ScheduleExceptionService from '../services/ScheduleExceptionService'

class ListScheduleExceptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                scheduleExceptions: []
        }
        this.addScheduleException = this.addScheduleException.bind(this);
        this.editScheduleException = this.editScheduleException.bind(this);
        this.deleteScheduleException = this.deleteScheduleException.bind(this);
    }

    deleteScheduleException(id){
        ScheduleExceptionService.deleteScheduleException(id).then( res => {
            this.setState({scheduleExceptions: this.state.scheduleExceptions.filter(scheduleException => scheduleException.scheduleExceptionId !== id)});
        });
    }
    viewScheduleException(id){
        this.props.history.push(`/view-scheduleException/${id}`);
    }
    editScheduleException(id){
        this.props.history.push(`/add-scheduleException/${id}`);
    }

    componentDidMount(){
        ScheduleExceptionService.getScheduleExceptions().then((res) => {
            this.setState({ scheduleExceptions: res.data});
        });
    }

    addScheduleException(){
        this.props.history.push('/add-scheduleException/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ScheduleException List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addScheduleException}> Add ScheduleException</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Date </th>
                                    <th> Reason </th>
                                    <th> Hours </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.scheduleExceptions.map(
                                        scheduleException => 
                                        <tr key = {scheduleException.scheduleExceptionId}>
                                             <td> { scheduleException.date } </td>
                                             <td> { scheduleException.reason } </td>
                                             <td> { scheduleException.hours } </td>
                                             <td>
                                                 <button onClick={ () => this.editScheduleException(scheduleException.scheduleExceptionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteScheduleException(scheduleException.scheduleExceptionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewScheduleException(scheduleException.scheduleExceptionId)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListScheduleExceptionComponent
