import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { FlightHealthEventService } from '../../../services/FlightHealthEvent.service';
import { FlightHealthEvent } from '../../../models/FlightHealthEvent';
import { SubBaseComponent } from '../../FlightHealthEvent/sub.base.component';

@Component({
    selector: 'app-create-flightHealthEvent',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateFlightHealthEventComponent extends SubBaseComponent implements OnInit {

    title = 'Add FlightHealthEvent';

    flightHealthEventForm: FormGroup;
    flightHealthEvent: FlightHealthEvent;

    constructor( http: HttpClient,
        private flightHealthEventService: FlightHealthEventService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.flightHealthEventForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  eventCode: ['', Validators.required],
      ConnectedAircraft: ['', ],
      Severity: ['', ]
        });
    }

    
    addFlightHealthEvent(eventCode, ConnectedAircraft, Severity): void {
        this.flightHealthEventService
        .addFlightHealthEvent(eventCode, ConnectedAircraft, Severity)
            .subscribe(() => {
                this.router.navigate(['/indexFlightHealthEvent']);
            });
    }

    ngOnInit(): void {
    }
}