import React, { Component } from 'react'
import CabinLayoutService from '../services/CabinLayoutService';

class CreateCabinLayoutComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                layoutCode: '',
                totalSeats: '',
                classLayout: ''
        }
        this.changelayoutCodeHandler = this.changelayoutCodeHandler.bind(this);
        this.changetotalSeatsHandler = this.changetotalSeatsHandler.bind(this);
        this.changeclassLayoutHandler = this.changeclassLayoutHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CabinLayoutService.getCabinLayoutById(this.state.id).then( (res) =>{
                let cabinLayout = res.data;
                this.setState({
                    layoutCode: cabinLayout.layoutCode,
                    totalSeats: cabinLayout.totalSeats,
                    classLayout: cabinLayout.classLayout
                });
            });
        }        
    }
    saveOrUpdateCabinLayout = (e) => {
        e.preventDefault();
        let cabinLayout = {
                cabinLayoutId: this.state.id,
                layoutCode: this.state.layoutCode,
                totalSeats: this.state.totalSeats,
                classLayout: this.state.classLayout
            };
        console.log('cabinLayout => ' + JSON.stringify(cabinLayout));

        // step 5
        if(this.state.id === '_add'){
            cabinLayout.cabinLayoutId=''
            CabinLayoutService.createCabinLayout(cabinLayout).then(res =>{
                this.props.history.push('/cabinLayouts');
            });
        }else{
            CabinLayoutService.updateCabinLayout(cabinLayout).then( res => {
                this.props.history.push('/cabinLayouts');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CabinLayout</h3>
        }else{
            return <h3 className="text-center">Update CabinLayout</h3>
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
                                            <label> layoutCode:&emsp; </label>
                                                <input placeholder="layoutCode" name="layoutCode" className="form-control" value={this.state.layoutCode} onChange={this.changelayoutCodeHandler}/>

                                            <label> totalSeats:&emsp; </label>
                                                <input type="number" placeholder="totalSeats" name="totalSeats" className="form-control" value={this.state.totalSeats} onChange={this.changetotalSeatsHandler}/>

                                            <label> classLayout:&emsp; </label>
                                                <input placeholder="classLayout" name="classLayout" className="form-control" value={this.state.classLayout} onChange={this.changeclassLayoutHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCabinLayout}>Save</button>
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

export default CreateCabinLayoutComponent
