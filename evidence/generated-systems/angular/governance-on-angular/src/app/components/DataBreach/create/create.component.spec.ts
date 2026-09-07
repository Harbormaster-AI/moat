
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateDataBreachComponent } from './create.component';
import { DataBreachService } from '../../../services/DataBreach.service';
import { Router } from '@angular/router';

describe('CreateDataBreachComponent', () => {
  let component: CreateDataBreachComponent;
  let fixture: ComponentFixture<CreateDataBreachComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateDataBreachComponent
      ],
      providers: [
        DataBreachService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateDataBreachComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});