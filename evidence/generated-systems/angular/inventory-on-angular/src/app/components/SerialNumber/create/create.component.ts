import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SerialNumberService } from '../../../services/SerialNumber.service';
import { SerialNumber } from '../../../models/SerialNumber';
import { SubBaseComponent } from '../../SerialNumber/sub.base.component';

@Component({
    selector: 'app-create-serialNumber',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateSerialNumberComponent extends SubBaseComponent implements OnInit {

    title = 'Add SerialNumber';

    serialNumberForm: FormGroup;
    serialNumber: SerialNumber;

    constructor( http: HttpClient,
        private serialNumberService: SerialNumberService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.serialNumberForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  serial: ['', Validators.required],
      activationDate: ['', Validators.required],
      Sku: ['', ],
      CurrentInventoryItem: ['', ],
      Lot: ['', ],
      Status: ['', ]
        });
    }

    
    addSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status): void {
        this.serialNumberService
        .addSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status)
            .subscribe(() => {
                this.router.navigate(['/indexSerialNumber']);
            });
    }

    ngOnInit(): void {
    }
}