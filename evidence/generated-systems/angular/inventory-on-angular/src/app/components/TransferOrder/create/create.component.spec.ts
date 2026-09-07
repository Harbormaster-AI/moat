
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateTransferOrderComponent } from './create.component';
import { TransferOrderService } from '../../../services/TransferOrder.service';
import { Router } from '@angular/router';

describe('CreateTransferOrderComponent', () => {
  let component: CreateTransferOrderComponent;
  let fixture: ComponentFixture<CreateTransferOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateTransferOrderComponent
      ],
      providers: [
        TransferOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateTransferOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});