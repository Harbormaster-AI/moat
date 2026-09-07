
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSalaryComponentComponent } from './create.component';
import { SalaryComponentService } from '../../../services/SalaryComponent.service';
import { Router } from '@angular/router';

describe('CreateSalaryComponentComponent', () => {
  let component: CreateSalaryComponentComponent;
  let fixture: ComponentFixture<CreateSalaryComponentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSalaryComponentComponent
      ],
      providers: [
        SalaryComponentService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSalaryComponentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});