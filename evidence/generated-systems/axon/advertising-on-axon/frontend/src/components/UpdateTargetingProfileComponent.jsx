import React, { Component } from 'react'
import TargetingProfileService from '../services/TargetingProfileService';

class UpdateTargetingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: ''
        }
        this.updateTargetingProfile = this.updateTargetingProfile.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        TargetingProfileService.getTargetingProfileById(this.state.id).then( (res) =>{
            let targetingProfile = res.data;
            this.setState({
                name: targetingProfile.name
            });
        });
    }

    updateTargetingProfile = (e) => {
        e.preventDefault();
        let targetingProfile = {
            targetingProfileId: this.state.id,
            name: this.state.name
        };
        console.log('targetingProfile => ' + JSON.stringify(targetingProfile));
        console.log('id => ' + JSON.stringify(this.state.id));
        TargetingProfileService.updateTargetingProfile(targetingProfile).then( res => {
            this.props.history.push('/targetingProfiles');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/targetingProfiles');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update TargetingProfile</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateTargetingProfile}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateTargetingProfileComponent
