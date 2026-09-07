import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { LotService } from '../../../services/Lot.service';
import { Lot } from '../../../models/Lot';
import { SubBaseComponent } from '../../Lot/sub.base.component';

@Component({
    selector: 'app-create-lot',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateLotComponent extends SubBaseComponent implements OnInit {

    title = 'Add Lot';

    lotForm: FormGroup;
    lot: Lot;

    constructor( http: HttpClient,
        private lotService: LotService,
        private fb: FormBuilder,
        private router: Router
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

    
    addLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus): void {
        this.lotService
        .addLot(batchNumber, manufactureDate, expirationDate, Sku, InventoryItems, LotStatus)
            .subscribe(() => {
                this.router.navigate(['/indexLot']);
            });
    }

    ngOnInit(): void {
    }
}