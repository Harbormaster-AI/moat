import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConnectedAircraftService } from '../../../services/ConnectedAircraft.service';
import { ConnectedAircraft } from '../../../models/ConnectedAircraft';
import { SubBaseComponent } from '../../ConnectedAircraft/sub.base.component';

@Component({
    selector: 'app-create-connectedAircraft',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateConnectedAircraftComponent extends SubBaseComponent implements OnInit {

    title = 'Add ConnectedAircraft';

    connectedAircraftForm: FormGroup;
    connectedAircraft: ConnectedAircraft;

    constructor( http: HttpClient,
        private connectedAircraftService: ConnectedAircraftService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.connectedAircraftForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  communicationsProvider: ['', Validators.required],
      Aircraft: ['', ],
      FlightHealthEvents: ['', ],
      SoftwareLoads: ['', ],
      ConnectivityStatus: ['', ]
        });
    }

    
    addConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus): void {
        this.connectedAircraftService
        .addConnectedAircraft(communicationsProvider, Aircraft, FlightHealthEvents, SoftwareLoads, ConnectivityStatus)
            .subscribe(() => {
                this.router.navigate(['/indexConnectedAircraft']);
            });
    }

    ngOnInit(): void {
    }
}