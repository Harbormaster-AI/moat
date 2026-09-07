
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateSoftwareUpdateComponent } from './create.component';
import { SoftwareUpdateService } from '../../../services/SoftwareUpdate.service';
import { Router } from '@angular/router';

describe('CreateSoftwareUpdateComponent', () => {
  let component: CreateSoftwareUpdateComponent;
  let fixture: ComponentFixture<CreateSoftwareUpdateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateSoftwareUpdateComponent
      ],
      providers: [
        SoftwareUpdateService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateSoftwareUpdateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});