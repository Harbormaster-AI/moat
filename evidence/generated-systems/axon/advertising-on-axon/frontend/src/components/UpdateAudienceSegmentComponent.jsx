import React, { Component } from 'react'
import AudienceSegmentService from '../services/AudienceSegmentService';

class UpdateAudienceSegmentComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                estimatedReach: '',
                description: '',
                providerType: ''
        }
        this.updateAudienceSegment = this.updateAudienceSegment.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeestimatedReachHandler = this.changeestimatedReachHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeProviderTypeHandler = this.changeProviderTypeHandler.bind(this);
    }

    componentDidMount(){
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

    updateAudienceSegment = (e) => {
        e.preventDefault();
        let audienceSegment = {
            audienceSegmentId: this.state.id,
            name: this.state.name,
            estimatedReach: this.state.estimatedReach,
            description: this.state.description,
            providerType: this.state.providerType
        };
        console.log('audienceSegment => ' + JSON.stringify(audienceSegment));
        console.log('id => ' + JSON.stringify(this.state.id));
        AudienceSegmentService.updateAudienceSegment(audienceSegment).then( res => {
            this.props.history.push('/audienceSegments');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AudienceSegment</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> estimatedReach: </label>
                                                <input type="number" placeholder="estimatedReach" name="estimatedReach" className="form-control" value={this.state.estimatedReach} onChange={this.changeestimatedReachHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> ProviderType: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAudienceSegment}>Save</button>
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

export default UpdateAudienceSegmentComponent
