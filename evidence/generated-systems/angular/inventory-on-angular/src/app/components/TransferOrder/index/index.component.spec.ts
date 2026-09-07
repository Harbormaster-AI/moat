
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTransferOrderComponent } from './index.component';
import { TransferOrderService } from '../../../services/TransferOrder.service';

describe('IndexTransferOrderComponent', () => {
  let component: IndexTransferOrderComponent;
  let fixture: ComponentFixture<IndexTransferOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTransferOrderComponent
      ],
      providers: [
        TransferOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTransferOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});