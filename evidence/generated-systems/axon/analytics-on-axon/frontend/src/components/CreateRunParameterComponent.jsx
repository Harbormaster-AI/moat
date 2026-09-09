import React, { Component } from 'react'
import RunParameterService from '../services/RunParameterService';

class CreateRunParameterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                value: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RunParameterService.getRunParameterById(this.state.id).then( (res) =>{
                let runParameter = res.data;
                this.setState({
                    name: runParameter.name,
                    value: runParameter.value
                });
            });
        }        
    }
    saveOrUpdateRunParameter = (e) => {
        e.preventDefault();
        let runParameter = {
                runParameterId: this.state.id,
                name: this.state.name,
                value: this.state.value
            };
        console.log('runParameter => ' + JSON.stringify(runParameter));

        // step 5
        if(this.state.id === '_add'){
            runParameter.runParameterId=''
            RunParameterService.createRunParameter(runParameter).then(res =>{
                this.props.history.push('/runParameters');
            });
        }else{
            RunParameterService.updateRunParameter(runParameter).then( res => {
                this.props.history.push('/runParameters');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changevalueHandler= (event) => {
        this.setState({value: event.target.value});
    }

    cancel(){
        this.props.history.push('/runParameters');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add RunParameter</h3>
        }else{
            return <h3 className="text-center">Update RunParameter</h3>
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

                                            <label> value:&emsp; </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRunParameter}>Save</button>
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

export default CreateRunParameterComponent
