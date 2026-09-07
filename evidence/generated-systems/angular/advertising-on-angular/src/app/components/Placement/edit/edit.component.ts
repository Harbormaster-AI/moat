import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { PlacementService } from '../../../services/Placement.service';
import { SubBaseComponent } from '../../Placement/sub.base.component';


@Component({
    selector: 'app-edit-placement',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditPlacementComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Placement';

    placementForm: FormGroup;
    placement: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: PlacementService,
        private fb: FormBuilder
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

    
    updatePlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal): void {
        this.route.params.subscribe((params) => {

                        this.service.updatePlacement(name, flight, goalImpressions, LineItem, AdSlot, Deal, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexPlacement']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getPlacement(params['id']).subscribe(res => {
                this.placement = res;
            });
        });
    }
}