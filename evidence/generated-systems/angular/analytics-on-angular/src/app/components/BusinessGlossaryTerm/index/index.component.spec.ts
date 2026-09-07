
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexBusinessGlossaryTermComponent } from './index.component';
import { BusinessGlossaryTermService } from '../../../services/BusinessGlossaryTerm.service';

describe('IndexBusinessGlossaryTermComponent', () => {
  let component: IndexBusinessGlossaryTermComponent;
  let fixture: ComponentFixture<IndexBusinessGlossaryTermComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexBusinessGlossaryTermComponent
      ],
      providers: [
        BusinessGlossaryTermService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexBusinessGlossaryTermComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});