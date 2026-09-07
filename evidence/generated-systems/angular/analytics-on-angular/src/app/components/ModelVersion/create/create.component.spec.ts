
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateModelVersionComponent } from './create.component';
import { ModelVersionService } from '../../../services/ModelVersion.service';
import { Router } from '@angular/router';

describe('CreateModelVersionComponent', () => {
  let component: CreateModelVersionComponent;
  let fixture: ComponentFixture<CreateModelVersionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateModelVersionComponent
      ],
      providers: [
        ModelVersionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateModelVersionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});