import React, { Component } from 'react'
import AudienceSegmentService from '../services/AudienceSegmentService';

class CreateAudienceSegmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                estimatedReach: '',
                description: '',
                providerType: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeestimatedReachHandler = this.changeestimatedReachHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AudienceSegmentService.getAudienceSegmentById(this.state.id).then( (res) =>{
                let audienceSegment = res.data;
                this.setState({
                    name: audienceSegment.name,
                    estimatedReach: audienceSegment.estimatedReach,
                    description: audienceSegment.description,
                    providerType: audienceSegment.providerType
                });
            });
        }        
    }
    saveOrUpdateAudienceSegment = (e) => {
        e.preventDefault();
        let audienceSegment = {
                audienceSegmentId: this.state.id,
                name: this.state.name,
                estimatedReach: this.state.estimatedReach,
                description: this.state.description,
                providerType: this.state.providerType
            };
        console.log('audienceSegment => ' + JSON.stringify(audienceSegment));

        // step 5
        if(this.state.id === '_add'){
            audienceSegment.audienceSegmentId=''
            AudienceSegmentService.createAudienceSegment(audienceSegment).then(res =>{
                this.props.history.push('/audienceSegments');
            });
        }else{
            AudienceSegmentService.updateAudienceSegment(audienceSegment).then( res => {
                this.props.history.push('/audienceSegments');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeestimatedReachHandler= (event) => {
        this.setState({estimatedReach: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeProviderTypeHandler= (event) => {
        this.setState({providerType: event.target.value});
    }

    cancel(){
        this.props.history.push('/audienceSegments');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AudienceSegment</h3>
        }else{
            return <h3 className="text-center">Update AudienceSegment</h3>
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

                                            <label> estimatedReach:&emsp; </label>
                                                <input type="number" placeholder="estimatedReach" name="estimatedReach" className="form-control" value={this.state.estimatedReach} onChange={this.changeestimatedReachHandler}/>

                                            <label> description:&emsp; </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> ProviderType:&emsp; </label>
                                                <select value={this.state.providerType} onChange={this.changeProviderTypeHandler}>
                      <option name="ProviderType" className="form-control" >
                          FirstParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          SecondParty
                      </option>
                      <option name="ProviderType" className="form-control" >
                          ThirdParty
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAudienceSegment}>Save</button>
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

export default CreateAudienceSegmentComponent
