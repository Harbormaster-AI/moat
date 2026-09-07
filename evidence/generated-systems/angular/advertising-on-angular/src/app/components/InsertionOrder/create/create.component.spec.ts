
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsertionOrderComponent } from './create.component';
import { InsertionOrderService } from '../../../services/InsertionOrder.service';
import { Router } from '@angular/router';

describe('CreateInsertionOrderComponent', () => {
  let component: CreateInsertionOrderComponent;
  let fixture: ComponentFixture<CreateInsertionOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsertionOrderComponent
      ],
      providers: [
        InsertionOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsertionOrderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});