
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateInsuredObjectComponent } from './create.component';
import { InsuredObjectService } from '../../../services/InsuredObject.service';
import { Router } from '@angular/router';

describe('CreateInsuredObjectComponent', () => {
  let component: CreateInsuredObjectComponent;
  let fixture: ComponentFixture<CreateInsuredObjectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateInsuredObjectComponent
      ],
      providers: [
        InsuredObjectService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateInsuredObjectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});