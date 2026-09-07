import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { PlacementService } from '../../../services/Placement.service';
import { Placement } from '../../../models/Placement';
import { SubBaseComponent } from '../../Placement/sub.base.component';

@Component({
    selector: 'app-create-placement',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreatePlacementComponent extends SubBaseComponent implements OnInit {

    title = 'Add Placement';

    placementForm: FormGroup;
    placement: Placement;

    constructor( http: HttpClient,
        private placementService: PlacementService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.placementForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      flight: ['', Validators.required],
      goalImpressions: ['', Validators.required],
      LineItem: ['', ],
      AdSlot: ['', ],
      Deal: ['', ]
        });
    }

    
    addPlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal): void {
        this.placementService
        .addPlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal)
            .subscribe(() => {
                this.router.navigate(['/indexPlacement']);
            });
    }

    ngOnInit(): void {
    }
}