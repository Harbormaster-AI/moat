import React, { Component } from 'react'
import CabinLayoutService from '../services/CabinLayoutService';

class UpdateCabinLayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                layoutCode: '',
                totalSeats: '',
                classLayout: ''
        }
        this.updateCabinLayout = this.updateCabinLayout.bind(this);

        this.changelayoutCodeHandler = this.changelayoutCodeHandler.bind(this);
        this.changetotalSeatsHandler = this.changetotalSeatsHandler.bind(this);
        this.changeclassLayoutHandler = this.changeclassLayoutHandler.bind(this);
    }

    componentDidMount(){
        CabinLayoutService.getCabinLayoutById(this.state.id).then( (res) =>{
            let cabinLayout = res.data;
            this.setState({
                layoutCode: cabinLayout.layoutCode,
                totalSeats: cabinLayout.totalSeats,
                classLayout: cabinLayout.classLayout
            });
        });
    }

    updateCabinLayout = (e) => {
        e.preventDefault();
        let cabinLayout = {
            cabinLayoutId: this.state.id,
            layoutCode: this.state.layoutCode,
            totalSeats: this.state.totalSeats,
            classLayout: this.state.classLayout
        };
        console.log('cabinLayout => ' + JSON.stringify(cabinLayout));
        console.log('id => ' + JSON.stringify(this.state.id));
        CabinLayoutService.updateCabinLayout(cabinLayout).then( res => {
            this.props.history.push('/cabinLayouts');
        });
    }

    changelayoutCodeHandler= (event) => {
        this.setState({layoutCode: event.target.value});
    }
    changetotalSeatsHandler= (event) => {
        this.setState({totalSeats: event.target.value});
    }
    changeclassLayoutHandler= (event) => {
        this.setState({classLayout: event.target.value});
    }

    cancel(){
        this.props.history.push('/cabinLayouts');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CabinLayout</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> layoutCode: </label>
                                                <input placeholder="layoutCode" name="layoutCode" className="form-control" value={this.state.layoutCode} onChange={this.changelayoutCodeHandler}/>

                                            <label> totalSeats: </label>
                                                <input type="number" placeholder="totalSeats" name="totalSeats" className="form-control" value={this.state.totalSeats} onChange={this.changetotalSeatsHandler}/>

                                            <label> classLayout: </label>
                                                <input placeholder="classLayout" name="classLayout" className="form-control" value={this.state.classLayout} onChange={this.changeclassLayoutHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCabinLayout}>Save</button>
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

export default UpdateCabinLayoutComponent
