import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { DemandSignalService } from '../../../services/DemandSignal.service';
import { DemandSignal } from '../../../models/DemandSignal';
import { SubBaseComponent } from '../../DemandSignal/sub.base.component';

@Component({
    selector: 'app-create-demandSignal',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateDemandSignalComponent extends SubBaseComponent implements OnInit {

    title = 'Add DemandSignal';

    demandSignalForm: FormGroup;
    demandSignal: DemandSignal;

    constructor( http: HttpClient,
        private demandSignalService: DemandSignalService,
        private fb: FormBuilder,
        private router: Router
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

    
    addDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType): void {
        this.demandSignalService
        .addDemandSignal(externalReference, requestedDate, quantity, Sku, Reservations, DemandType)
            .subscribe(() => {
                this.router.navigate(['/indexDemandSignal']);
            });
    }

    ngOnInit(): void {
    }
}