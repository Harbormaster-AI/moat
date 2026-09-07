
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateLineItemComponent } from './create.component';
import { LineItemService } from '../../../services/LineItem.service';
import { Router } from '@angular/router';

describe('CreateLineItemComponent', () => {
  let component: CreateLineItemComponent;
  let fixture: ComponentFixture<CreateLineItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateLineItemComponent
      ],
      providers: [
        LineItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateLineItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});