import React, { Component } from 'react'
import DataProcessingActivityService from '../services/DataProcessingActivityService';

class CreateDataProcessingActivityComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                purpose: '',
                startDate: '',
                lawfulBasis: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changepurposeHandler = this.changepurposeHandler.bind(this);
        this.changestartDateHandler = this.changestartDateHandler.bind(this);
        this.changeLawfulBasisHandler = this.changeLawfulBasisHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateDataProcessingActivity = (e) => {
        e.preventDefault();
        let dataProcessingActivity = {
                dataProcessingActivityId: this.state.id,
                name: this.state.name,
                purpose: this.state.purpose,
                startDate: this.state.startDate,
                lawfulBasis: this.state.lawfulBasis
            };
        console.log('dataProcessingActivity => ' + JSON.stringify(dataProcessingActivity));

        // step 5
        if(this.state.id === '_add'){
            dataProcessingActivity.dataProcessingActivityId=''
            DataProcessingActivityService.createDataProcessingActivity(dataProcessingActivity).then(res =>{
                this.props.history.push('/dataProcessingActivitys');
            });
        }else{
            DataProcessingActivityService.updateDataProcessingActivity(dataProcessingActivity).then( res => {
                this.props.history.push('/dataProcessingActivitys');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add DataProcessingActivity</h3>
        }else{
            return <h3 className="text-center">Update DataProcessingActivity</h3>
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

                                            <label> purpose:&emsp; </label>
                                                <input placeholder="purpose" name="purpose" className="form-control" value={this.state.purpose} onChange={this.changepurposeHandler}/>

                                            <label> startDate:&emsp; </label>
                                                <input type="date" placeholder="startDate" name="startDate" className="form-control" value={this.state.startDate} onChange={this.changestartDateHandler}/>

                                            <label> LawfulBasis:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateDataProcessingActivity}>Save</button>
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

export default CreateDataProcessingActivityComponent
