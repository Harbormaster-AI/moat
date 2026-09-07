import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { LotService } from '../../../services/Lot.service';
import { SubBaseComponent } from '../../Lot/sub.base.component';


@Component({
    selector: 'app-edit-lot',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditLotComponent extends SubBaseComponent implements OnInit {

    title = 'Edit Lot';

    lotForm: FormGroup;
    lot: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: LotService,
        private fb: FormBuilder
) {
        super(http);
        this.lotForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  batchNumber: ['', Validators.required],
      manufactureDate: ['', Validators.required],
      expirationDate: ['', Validators.required],
      Sku: ['', ],
      InventoryItems: ['', ],
      LotStatus: ['', ]
        });
    }

    
    updateLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus): void {
        this.route.params.subscribe((params) => {

                        this.service.updateLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexLot']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getLot(params['id']).subscribe(res => {
                this.lot = res;
            });
        });
    }
}