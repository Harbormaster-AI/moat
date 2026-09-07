
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { CreateBusinessGlossaryTermComponent } from './create.component';
import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';
import { Router } from '@angular/router';

describe('CreateBusinessGlossaryTermComponent', () => {
  let component: CreateBusinessGlossaryTermComponent;
  let fixture: ComponentFixture<CreateBusinessGlossaryTermComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        CreateBusinessGlossaryTermComponent
      ],
      providers: [
        BusinessGlossaryTermService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(CreateBusinessGlossaryTermComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});