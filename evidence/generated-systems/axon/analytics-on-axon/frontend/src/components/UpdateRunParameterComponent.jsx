import React, { Component } from 'react'
import RunParameterService from '../services/RunParameterService';

class UpdateRunParameterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                value: ''
        }
        this.updateRunParameter = this.updateRunParameter.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changevalueHandler = this.changevalueHandler.bind(this);
    }

    componentDidMount(){
        RunParameterService.getRunParameterById(this.state.id).then( (res) =>{
            let runParameter = res.data;
            this.setState({
                name: runParameter.name,
                value: runParameter.value
            });
        });
    }

    updateRunParameter = (e) => {
        e.preventDefault();
        let runParameter = {
            runParameterId: this.state.id,
            name: this.state.name,
            value: this.state.value
        };
        console.log('runParameter => ' + JSON.stringify(runParameter));
        console.log('id => ' + JSON.stringify(this.state.id));
        RunParameterService.updateRunParameter(runParameter).then( res => {
            this.props.history.push('/runParameters');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update RunParameter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> value: </label>
                                                <input placeholder="value" name="value" className="form-control" value={this.state.value} onChange={this.changevalueHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRunParameter}>Save</button>
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

export default UpdateRunParameterComponent
