import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { SerialNumberService } from '../../../services/SerialNumber.service';
import { SubBaseComponent } from '../../SerialNumber/sub.base.component';


@Component({
    selector: 'app-edit-serialNumber',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditSerialNumberComponent extends SubBaseComponent implements OnInit {

    title = 'Edit SerialNumber';

    serialNumberForm: FormGroup;
    serialNumber: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: SerialNumberService,
        private fb: FormBuilder
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

    
    updateSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status): void {
        this.route.params.subscribe((params) => {

                        this.service.updateSerialNumber(serial, activationDate, Sku, CurrentInventoryItem, Lot, Status, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexSerialNumber']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getSerialNumber(params['id']).subscribe(res => {
                this.serialNumber = res;
            });
        });
    }
}