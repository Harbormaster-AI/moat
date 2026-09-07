
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateCoverageComponent } from './create.component';
import { CoverageService } from '../../../services/Coverage.service';
import { Router } from '@angular/router';

describe('CreateCoverageComponent', () => {
  let component: CreateCoverageComponent;
  let fixture: ComponentFixture<CreateCoverageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateCoverageComponent
      ],
      providers: [
        CoverageService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateCoverageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});