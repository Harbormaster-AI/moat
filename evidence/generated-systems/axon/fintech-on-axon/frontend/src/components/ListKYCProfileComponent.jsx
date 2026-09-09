import React, { Component } from 'react'
import KYCProfileService from '../services/KYCProfileService'

class ListKYCProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                kYCProfiles: []
        }
        this.addKYCProfile = this.addKYCProfile.bind(this);
        this.editKYCProfile = this.editKYCProfile.bind(this);
        this.deleteKYCProfile = this.deleteKYCProfile.bind(this);
    }

    deleteKYCProfile(id){
        KYCProfileService.deleteKYCProfile(id).then( res => {
            this.setState({kYCProfiles: this.state.kYCProfiles.filter(kYCProfile => kYCProfile.kYCProfileId !== id)});
        });
    }
    viewKYCProfile(id){
        this.props.history.push(`/view-kYCProfile/${id}`);
    }
    editKYCProfile(id){
        this.props.history.push(`/add-kYCProfile/${id}`);
    }

    componentDidMount(){
        KYCProfileService.getKYCProfiles().then((res) => {
            this.setState({ kYCProfiles: res.data});
        });
    }

    addKYCProfile(){
        this.props.history.push('/add-kYCProfile/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">KYCProfile List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addKYCProfile}> Add KYCProfile</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ProfileId </th>
                                    <th> CreatedAt </th>
                                    <th> Status </th>
                                    <th> VerificationLevel </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.kYCProfiles.map(
                                        kYCProfile => 
                                        <tr key = {kYCProfile.kYCProfileId}>
                                             <td> { kYCProfile.profileId } </td>
                                             <td> { kYCProfile.createdAt } </td>
                                             <td> { kYCProfile.status } </td>
                                             <td> { kYCProfile.verificationLevel } </td>
                                             <td>
                                                 <button onClick={ () => this.editKYCProfile(kYCProfile.kYCProfileId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteKYCProfile(kYCProfile.kYCProfileId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewKYCProfile(kYCProfile.kYCProfileId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListKYCProfileComponent
