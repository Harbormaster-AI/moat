import React, { Component } from 'react'
import ExperimentVariantService from '../services/ExperimentVariantService';

class UpdateExperimentVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                allocation: ''
        }
        this.updateExperimentVariant = this.updateExperimentVariant.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeallocationHandler = this.changeallocationHandler.bind(this);
    }

    componentDidMount(){
        ExperimentVariantService.getExperimentVariantById(this.state.id).then( (res) =>{
            let experimentVariant = res.data;
            this.setState({
                name: experimentVariant.name,
                allocation: experimentVariant.allocation
            });
        });
    }

    updateExperimentVariant = (e) => {
        e.preventDefault();
        let experimentVariant = {
            experimentVariantId: this.state.id,
            name: this.state.name,
            allocation: this.state.allocation
        };
        console.log('experimentVariant => ' + JSON.stringify(experimentVariant));
        console.log('id => ' + JSON.stringify(this.state.id));
        ExperimentVariantService.updateExperimentVariant(experimentVariant).then( res => {
            this.props.history.push('/experimentVariants');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeallocationHandler= (event) => {
        this.setState({allocation: event.target.value});
    }

    cancel(){
        this.props.history.push('/experimentVariants');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ExperimentVariant</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> allocation: </label>
                                                <input placeholder="allocation" name="allocation" className="form-control" value={this.state.allocation} onChange={this.changeallocationHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateExperimentVariant}>Save</button>
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

export default UpdateExperimentVariantComponent
