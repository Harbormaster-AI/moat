import React, { Component } from 'react'
import DataProcessingActivityService from '../services/DataProcessingActivityService';

class UpdateDataProcessingActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                purpose: '',
                startDate: '',
                lawfulBasis: ''
        }
        this.updateDataProcessingActivity = this.updateDataProcessingActivity.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changepurposeHandler = this.changepurposeHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeLawfulBasisHandler = this.changeLawfulBasisHandler.bind(this);
    }

    componentDidMount(){
        DataProcessingActivityService.getDataProcessingActivityById(this.state.id).then( (res) =>{
            let dataProcessingActivity = res.data;
            this.setState({
                name: dataProcessingActivity.name,
                purpose: dataProcessingActivity.purpose,
                startDate: dataProcessingActivity.startDate,
                lawfulBasis: dataProcessingActivity.lawfulBasis
            });
        });
    }

    updateDataProcessingActivity = (e) => {
        e.preventDefault();
        let dataProcessingActivity = {
            dataProcessingActivityId: this.state.id,
            name: this.state.name,
            purpose: this.state.purpose,
            startDate: this.state.startDate,
            lawfulBasis: this.state.lawfulBasis
        };
        console.log('dataProcessingActivity => ' + JSON.stringify(dataProcessingActivity));
        console.log('id => ' + JSON.stringify(this.state.id));
        DataProcessingActivityService.updateDataProcessingActivity(dataProcessingActivity).then( res => {
            this.props.history.push('/dataProcessingActivitys');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changepurposeHandler= (event) => {
        this.setState({purpose: event.target.value});
    }
    changestartDateHandler= (event) => {
        this.setState({startDate: event.target.value});
    }
    changeLawfulBasisHandler= (event) => {
        this.setState({lawfulBasis: event.target.value});
    }

    cancel(){
        this.props.history.push('/dataProcessingActivitys');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DataProcessingActivity</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> purpose: </label>
                                                <input placeholder="purpose" name="purpose" className="form-control" value={this.state.purpose} onChange={this.changepurposeHandler}/>

                                            <label> startDate: </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> LawfulBasis: </label>
                                                <select value={this.state.lawfulBasis} onChange={this.changeLawfulBasisHandler}>
                      <option name="LawfulBasis" className="form-control" >
                          Consent
                      </option>
                      <option name="LawfulBasis" className="form-control" >
                          Contract
                      </option>
                      <option name="LawfulBasis" className="form-control" >
                          LegalObligation
                      </option>
                      <option name="LawfulBasis" className="form-control" >
                          VitalInterests
                      </option>
                      <option name="LawfulBasis" className="form-control" >
                          PublicTask
                      </option>
                      <option name="LawfulBasis" className="form-control" >
                          LegitimateInterests
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDataProcessingActivity}>Save</button>
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

export default UpdateDataProcessingActivityComponent
