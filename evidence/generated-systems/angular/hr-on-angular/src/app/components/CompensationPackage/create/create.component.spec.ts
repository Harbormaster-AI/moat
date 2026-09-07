
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCompensationPackageComponent } from './create.component';
import { CompensationPackageService } from '../../../services/CompensationPackage.service';
import { Router } from '@angular/router';

describe('CreateCompensationPackageComponent', () => {
  let component: CreateCompensationPackageComponent;
  let fixture: ComponentFixture<CreateCompensationPackageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCompensationPackageComponent
      ],
      providers: [
        CompensationPackageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCompensationPackageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});