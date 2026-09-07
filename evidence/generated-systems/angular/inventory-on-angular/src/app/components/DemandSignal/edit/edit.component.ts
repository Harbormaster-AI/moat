import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { DemandSignalService } from '../../../services/DemandSignal.service';
import { SubBaseComponent } from '../../DemandSignal/sub.base.component';


@Component({
    selector: 'app-edit-demandSignal',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditDemandSignalComponent extends SubBaseComponent implements OnInit {

    title = 'Edit DemandSignal';

    demandSignalForm: FormGroup;
    demandSignal: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: DemandSignalService,
        private fb: FormBuilder
) {
        super(http);
        this.demandSignalForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  externalReference: ['', Validators.required],
      requestedDate: ['', Validators.required],
      quantity: ['', Validators.required],
      Sku: ['', ],
      Reservations: ['', ],
      DemandType: ['', ]
        });
    }

    
    updateDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexDemandSignal']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getDemandSignal(params['id']).subscribe(res => {
                this.demandSignal = res;
            });
        });
    }
}