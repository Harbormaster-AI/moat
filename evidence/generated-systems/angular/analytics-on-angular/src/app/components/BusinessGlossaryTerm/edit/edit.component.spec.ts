
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditBusinessGlossaryTermComponent } from './edit.component';
import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';

describe('EditBusinessGlossaryTermComponent', () => {
  let component: EditBusinessGlossaryTermComponent;
  let fixture: ComponentFixture<EditBusinessGlossaryTermComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditBusinessGlossaryTermComponent
      ],
      providers: [
        BusinessGlossaryTermService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditBusinessGlossaryTermComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});