
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSemanticModelComponent } from './index.component';
import { SemanticModelService } from '../../../services/SemanticModel.service';

describe('IndexSemanticModelComponent', () => {
  let component: IndexSemanticModelComponent;
  let fixture: ComponentFixture<IndexSemanticModelComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSemanticModelComponent
      ],
      providers: [
        SemanticModelService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSemanticModelComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});