import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { InventorySourceService } from '../../../services/InventorySource.service';
import { InventorySource } from '../../../models/InventorySource';
import { SubBaseComponent } from '../../InventorySource/sub.base.component';

@Component({
    selector: 'app-create-inventorySource',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateInventorySourceComponent extends SubBaseComponent implements OnInit {

    title = 'Add InventorySource';

    inventorySourceForm: FormGroup;
    inventorySource: InventorySource;

    constructor( http: HttpClient,
        private inventorySourceService: InventorySourceService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.inventorySourceForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      domain: ['', Validators.required],
      Publisher: ['', ],
      AdSlots: ['', ],
      Deals: ['', ],
      Channel: ['', ],
      PrimaryFormat: ['', ]
        });
    }

    
    addInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat): void {
        this.inventorySourceService
        .addInventorySource(name, domain, Publisher, AdSlots, Deals, Channel, PrimaryFormat)
            .subscribe(() => {
                this.router.navigate(['/indexInventorySource']);
            });
    }

    ngOnInit(): void {
    }
}