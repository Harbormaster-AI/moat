import React, { Component } from 'react'
import TargetingProfileService from '../services/TargetingProfileService';

class CreateTargetingProfileComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            TargetingProfileService.getTargetingProfileById(this.state.id).then( (res) =>{
                let targetingProfile = res.data;
                this.setState({
                    name: targetingProfile.name
                });
            });
        }        
    }
    saveOrUpdateTargetingProfile = (e) => {
        e.preventDefault();
        let targetingProfile = {
                targetingProfileId: this.state.id,
                name: this.state.name
            };
        console.log('targetingProfile => ' + JSON.stringify(targetingProfile));

        // step 5
        if(this.state.id === '_add'){
            targetingProfile.targetingProfileId=''
            TargetingProfileService.createTargetingProfile(targetingProfile).then(res =>{
                this.props.history.push('/targetingProfiles');
            });
        }else{
            TargetingProfileService.updateTargetingProfile(targetingProfile).then( res => {
                this.props.history.push('/targetingProfiles');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/targetingProfiles');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add TargetingProfile</h3>
        }else{
            return <h3 className="text-center">Update TargetingProfile</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateTargetingProfile}>Save</button>
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

export default CreateTargetingProfileComponent
