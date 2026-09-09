import React, { Component } from 'react'
import ExperimentVariantService from '../services/ExperimentVariantService';

class CreateExperimentVariantComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                allocation: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeallocationHandler = this.changeallocationHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ExperimentVariantService.getExperimentVariantById(this.state.id).then( (res) =>{
                let experimentVariant = res.data;
                this.setState({
                    name: experimentVariant.name,
                    allocation: experimentVariant.allocation
                });
            });
        }        
    }
    saveOrUpdateExperimentVariant = (e) => {
        e.preventDefault();
        let experimentVariant = {
                experimentVariantId: this.state.id,
                name: this.state.name,
                allocation: this.state.allocation
            };
        console.log('experimentVariant => ' + JSON.stringify(experimentVariant));

        // step 5
        if(this.state.id === '_add'){
            experimentVariant.experimentVariantId=''
            ExperimentVariantService.createExperimentVariant(experimentVariant).then(res =>{
                this.props.history.push('/experimentVariants');
            });
        }else{
            ExperimentVariantService.updateExperimentVariant(experimentVariant).then( res => {
                this.props.history.push('/experimentVariants');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ExperimentVariant</h3>
        }else{
            return <h3 className="text-center">Update ExperimentVariant</h3>
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

                                            <label> allocation:&emsp; </label>
                                                <input placeholder="allocation" name="allocation" className="form-control" value={this.state.allocation} onChange={this.changeallocationHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateExperimentVariant}>Save</button>
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

export default CreateExperimentVariantComponent
