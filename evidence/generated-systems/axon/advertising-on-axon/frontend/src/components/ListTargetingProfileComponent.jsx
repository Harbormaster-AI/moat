import React, { Component } from 'react'
import TargetingProfileService from '../services/TargetingProfileService'

class ListTargetingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                targetingProfiles: []
        }
        this.addTargetingProfile = this.addTargetingProfile.bind(this);
        this.editTargetingProfile = this.editTargetingProfile.bind(this);
        this.deleteTargetingProfile = this.deleteTargetingProfile.bind(this);
    }

    deleteTargetingProfile(id){
        TargetingProfileService.deleteTargetingProfile(id).then( res => {
            this.setState({targetingProfiles: this.state.targetingProfiles.filter(targetingProfile => targetingProfile.targetingProfileId !== id)});
        });
    }
    viewTargetingProfile(id){
        this.props.history.push(`/view-targetingProfile/${id}`);
    }
    editTargetingProfile(id){
        this.props.history.push(`/add-targetingProfile/${id}`);
    }

    componentDidMount(){
        TargetingProfileService.getTargetingProfiles().then((res) => {
            this.setState({ targetingProfiles: res.data});
        });
    }

    addTargetingProfile(){
        this.props.history.push('/add-targetingProfile/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">TargetingProfile List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addTargetingProfile}> Add TargetingProfile</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.targetingProfiles.map(
                                        targetingProfile => 
                                        <tr key = {targetingProfile.targetingProfileId}>
                                             <td> { targetingProfile.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editTargetingProfile(targetingProfile.targetingProfileId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteTargetingProfile(targetingProfile.targetingProfileId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewTargetingProfile(targetingProfile.targetingProfileId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListTargetingProfileComponent
