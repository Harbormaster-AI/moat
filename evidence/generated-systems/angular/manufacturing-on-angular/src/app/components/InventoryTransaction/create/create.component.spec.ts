
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInventoryTransactionComponent } from './create.component';
import { InventoryTransactionService } from '../../../services/InventoryTransaction.service';
import { Router } from '@angular/router';

describe('CreateInventoryTransactionComponent', () => {
  let component: CreateInventoryTransactionComponent;
  let fixture: ComponentFixture<CreateInventoryTransactionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInventoryTransactionComponent
      ],
      providers: [
        InventoryTransactionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInventoryTransactionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});