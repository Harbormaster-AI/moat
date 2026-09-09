import React, { Component } from 'react'
import AircraftOptionService from '../services/AircraftOptionService'

class ListAircraftOptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                aircraftOptions: []
        }
        this.addAircraftOption = this.addAircraftOption.bind(this);
        this.editAircraftOption = this.editAircraftOption.bind(this);
        this.deleteAircraftOption = this.deleteAircraftOption.bind(this);
    }

    deleteAircraftOption(id){
        AircraftOptionService.deleteAircraftOption(id).then( res => {
            this.setState({aircraftOptions: this.state.aircraftOptions.filter(aircraftOption => aircraftOption.aircraftOptionId !== id)});
        });
    }
    viewAircraftOption(id){
        this.props.history.push(`/view-aircraftOption/${id}`);
    }
    editAircraftOption(id){
        this.props.history.push(`/add-aircraftOption/${id}`);
    }

    componentDidMount(){
        AircraftOptionService.getAircraftOptions().then((res) => {
            this.setState({ aircraftOptions: res.data});
        });
    }

    addAircraftOption(){
        this.props.history.push('/add-aircraftOption/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">AircraftOption List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addAircraftOption}> Add AircraftOption</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Code </th>
                                    <th> Name </th>
                                    <th> OptionCategory </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.aircraftOptions.map(
                                        aircraftOption => 
                                        <tr key = {aircraftOption.aircraftOptionId}>
                                             <td> { aircraftOption.code } </td>
                                             <td> { aircraftOption.name } </td>
                                             <td> { aircraftOption.optionCategory } </td>
                                             <td>
                                                 <button onClick={ () => this.editAircraftOption(aircraftOption.aircraftOptionId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteAircraftOption(aircraftOption.aircraftOptionId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewAircraftOption(aircraftOption.aircraftOptionId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListAircraftOptionComponent
