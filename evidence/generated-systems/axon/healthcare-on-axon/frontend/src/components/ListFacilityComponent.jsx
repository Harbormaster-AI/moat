import React, { Component } from 'react'
import FacilityService from '../services/FacilityService'

class ListFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                facilitys: []
        }
        this.addFacility = this.addFacility.bind(this);
        this.editFacility = this.editFacility.bind(this);
        this.deleteFacility = this.deleteFacility.bind(this);
    }

    deleteFacility(id){
        FacilityService.deleteFacility(id).then( res => {
            this.setState({facilitys: this.state.facilitys.filter(facility => facility.facilityId !== id)});
        });
    }
    viewFacility(id){
        this.props.history.push(`/view-facility/${id}`);
    }
    editFacility(id){
        this.props.history.push(`/add-facility/${id}`);
    }

    componentDidMount(){
        FacilityService.getFacilitys().then((res) => {
            this.setState({ facilitys: res.data});
        });
    }

    addFacility(){
        this.props.history.push('/add-facility/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Facility List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addFacility}> Add Facility</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> FacilityCode </th>
                                    <th> Address </th>
                                    <th> FacilityType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.facilitys.map(
                                        facility => 
                                        <tr key = {facility.facilityId}>
                                             <td> { facility.name } </td>
                                             <td> { facility.facilityCode } </td>
                                             <td> { facility.address } </td>
                                             <td> { facility.facilityType } </td>
                                             <td>
                                                 <button onClick={ () => this.editFacility(facility.facilityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteFacility(facility.facilityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewFacility(facility.facilityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListFacilityComponent
