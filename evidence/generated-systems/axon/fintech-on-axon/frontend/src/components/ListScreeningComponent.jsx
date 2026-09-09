import React, { Component } from 'react'
import ScreeningService from '../services/ScreeningService'

class ListScreeningComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                screenings: []
        }
        this.addScreening = this.addScreening.bind(this);
        this.editScreening = this.editScreening.bind(this);
        this.deleteScreening = this.deleteScreening.bind(this);
    }

    deleteScreening(id){
        ScreeningService.deleteScreening(id).then( res => {
            this.setState({screenings: this.state.screenings.filter(screening => screening.screeningId !== id)});
        });
    }
    viewScreening(id){
        this.props.history.push(`/view-screening/${id}`);
    }
    editScreening(id){
        this.props.history.push(`/add-screening/${id}`);
    }

    componentDidMount(){
        ScreeningService.getScreenings().then((res) => {
            this.setState({ screenings: res.data});
        });
    }

    addScreening(){
        this.props.history.push('/add-screening/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Screening List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addScreening}> Add Screening</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Score </th>
                                    <th> ScreenedAt </th>
                                    <th> ScreeningType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.screenings.map(
                                        screening => 
                                        <tr key = {screening.screeningId}>
                                             <td> { screening.score } </td>
                                             <td> { screening.screenedAt } </td>
                                             <td> { screening.screeningType } </td>
                                             <td> { screening.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editScreening(screening.screeningId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteScreening(screening.screeningId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewScreening(screening.screeningId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListScreeningComponent
