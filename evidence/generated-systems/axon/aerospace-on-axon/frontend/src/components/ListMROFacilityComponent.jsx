import React, { Component } from 'react'
import MROFacilityService from '../services/MROFacilityService'

class ListMROFacilityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                mROFacilitys: []
        }
        this.addMROFacility = this.addMROFacility.bind(this);
        this.editMROFacility = this.editMROFacility.bind(this);
        this.deleteMROFacility = this.deleteMROFacility.bind(this);
    }

    deleteMROFacility(id){
        MROFacilityService.deleteMROFacility(id).then( res => {
            this.setState({mROFacilitys: this.state.mROFacilitys.filter(mROFacility => mROFacility.mROFacilityId !== id)});
        });
    }
    viewMROFacility(id){
        this.props.history.push(`/view-mROFacility/${id}`);
    }
    editMROFacility(id){
        this.props.history.push(`/add-mROFacility/${id}`);
    }

    componentDidMount(){
        MROFacilityService.getMROFacilitys().then((res) => {
            this.setState({ mROFacilitys: res.data});
        });
    }

    addMROFacility(){
        this.props.history.push('/add-mROFacility/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">MROFacility List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addMROFacility}> Add MROFacility</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> ApprovalScope </th>
                                    <th> Address </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.mROFacilitys.map(
                                        mROFacility => 
                                        <tr key = {mROFacility.mROFacilityId}>
                                             <td> { mROFacility.name } </td>
                                             <td> { mROFacility.approvalScope } </td>
                                             <td> { mROFacility.address } </td>
                                             <td>
                                                 <button onClick={ () => this.editMROFacility(mROFacility.mROFacilityId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteMROFacility(mROFacility.mROFacilityId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewMROFacility(mROFacility.mROFacilityId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListMROFacilityComponent
