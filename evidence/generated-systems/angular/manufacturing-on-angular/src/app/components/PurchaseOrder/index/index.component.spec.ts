
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPurchaseOrderComponent } from './index.component';
import { PurchaseOrderService } from '../../../services/PurchaseOrder.service';

describe('IndexPurchaseOrderComponent', () => {
  let component: IndexPurchaseOrderComponent;
  let fixture: ComponentFixture<IndexPurchaseOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPurchaseOrderComponent
      ],
      providers: [
        PurchaseOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPurchaseOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});