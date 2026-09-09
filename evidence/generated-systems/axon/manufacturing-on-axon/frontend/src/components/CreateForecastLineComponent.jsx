import React, { Component } from 'react'
import ForecastLineService from '../services/ForecastLineService';

class CreateForecastLineComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                period: '',
                quantity: '',
                confidence: ''
        }
        this.changeperiodHandler = this.changeperiodHandler.bind(this);
        this.changequantityHandler = this.changequantityHandler.bind(this);
        this.changeconfidenceHandler = this.changeconfidenceHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            ForecastLineService.getForecastLineById(this.state.id).then( (res) =>{
                let forecastLine = res.data;
                this.setState({
                    period: forecastLine.period,
                    quantity: forecastLine.quantity,
                    confidence: forecastLine.confidence
                });
            });
        }        
    }
    saveOrUpdateForecastLine = (e) => {
        e.preventDefault();
        let forecastLine = {
                forecastLineId: this.state.id,
                period: this.state.period,
                quantity: this.state.quantity,
                confidence: this.state.confidence
            };
        console.log('forecastLine => ' + JSON.stringify(forecastLine));

        // step 5
        if(this.state.id === '_add'){
            forecastLine.forecastLineId=''
            ForecastLineService.createForecastLine(forecastLine).then(res =>{
                this.props.history.push('/forecastLines');
            });
        }else{
            ForecastLineService.updateForecastLine(forecastLine).then( res => {
                this.props.history.push('/forecastLines');
            });
        }
    }
    
    changeperiodHandler= (event) => {
        this.setState({period: event.target.value});
    }
    changequantityHandler= (event) => {
        this.setState({quantity: event.target.value});
    }
    changeconfidenceHandler= (event) => {
        this.setState({confidence: event.target.value});
    }

    cancel(){
        this.props.history.push('/forecastLines');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add ForecastLine</h3>
        }else{
            return <h3 className="text-center">Update ForecastLine</h3>
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
                                            <label> period:&emsp; </label>
                                                <input type="date" placeholder="period" name="period" className="form-control" value={this.state.period} onChange={this.changeperiodHandler}/>

                                            <label> quantity:&emsp; </label>
                                                <input placeholder="quantity" name="quantity" className="form-control" value={this.state.quantity} onChange={this.changequantityHandler}/>

                                            <label> confidence:&emsp; </label>
                                                <input placeholder="confidence" name="confidence" className="form-control" value={this.state.confidence} onChange={this.changeconfidenceHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateForecastLine}>Save</button>
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

export default CreateForecastLineComponent
